# DatabaseLayer_Server 

Go REST API with CRUD operations for Knox database
The code can be found in this [repository](https://github.com/Knox-AAU/DatabaseLayer_Server).

The Access API documentation can temporarily be found [here](/DataLayer/access-api)

## Branch protection rules

The main branch of the repository is protected by branch protection rules, which ensures the code can build and is tested before it can be merged to the main branch. The branch protection rule requires pull requests before merging and status checks to pass before merging. The status checks are defined in the `test.yml` file from the Test workflow in GitHub Actions.

## Generate new documentation based on code

### Prerequisites

- We use `go-swagger` to generate Open-API 2.0 spec files automatically from code. See [go-swagger](https://goswagger.io) for more info.
- To turn the spec file into markdown, we use [openapi-markdown](https://github.com/theBenForce/openapi-markdown).

### Generate new documentation

To generate new documentation, run:

```bash
swagger generate spec -m -o ./swagger.yaml
```

To convert the spec file to markdown, run:

```bash
openapi-markdown -i ./swagger.yaml
```

Remember to insert the new markdown file in the wiki manually using `ctrl+a ctrl+c ctrl+v` on windows.

## Servers and access

To access the servers, you must either be on AAU's network or be using the [AAU VPN](https://www.en.its.aau.dk/instructions/vpn), and have permissions from ITS.

The servers are:

- `knox-kb01.srv.aau.dk` (knowledge graph databases)  
- `knox-db01.srv.aau.dk` (relational databases)  
- `knox-func01.srv.aau.dk` (functional layer)  
- `knox-preproc01.srv.aau.dk` (preprocessing layer)  
- `knox-web01.srv.aau.dk` (front-end services)  
- `knox-front01.srv.aau.dk` (front-end proxy)  
- `knox-proxy01.srv.aau.dk` (API proxy)  

You can access the server from the command line with `ssh <STUDENT_MAIL>@<SERVER_NAME> -L <your_port>:localhost:<host_port>`.

## Accessing our server and API

### Access our Virtuoso database

The code for this project (including tests) accesses Virtuoso on port 8890, which is the same port to access it on the server.

`ssh <STUDENT_MAIL>@knox-kb01.srv.aau.dk -L 8890:localhost:8890`

### Access our database layer API

`ssh <STUDENT_MAIL>@knox-kb01.srv.aau.dk -L <your_port>:localhost:80`

#### Deploy new version manually

Deployment is normally handled by watchtower on push to main. However, in case of the need of manual deployment, run

```bash
docker run -p 0.0.0.0:80:8000 --add-host=host.docker.internal:host-gateway -e VIRTUOSO_SERVER_URL=http://host.docker.internal:8890/sparql/ -e VIRTUOSO_GRAPH_URI=http://knox_ontology/ -e VIRTUOSO_ONTOLOGY_GRAPH_URI=http://knox_ontology/ -e VIRTUOSO_TEST_GRAPH_URI=http://testing/ -e VIRTUOSO_USERNAME=dba -e VIRTUOSO_PASSWORD=*** -d ghcr.io/knox-aau/databaselayer_server:main
```

Note that the ports map to the ports used in the ssh command. 

## Endpoint Documentation

### /get

#### GET
##### Summary:

This endpoint allows for querying with filters.

##### Description:

Example query: {{url}}/get?p=x&p=y&s=x&s=y&o=x&o=y

To query the whole graph, leave all parameters empty.

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| s | query | Subjects | No | [ string ] |
| o | query | Objects | No | [ string ] |
| p | query | Predicates | No | [ string ] |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | filtered triples response | [ [Result](#result) ] |

### Models


#### BindingAttribute

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| Type | string |  | No |
| Value | string |  | No |

#### Result

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| query | string |  | No |
| triples | [ [Triple](#triple) ] |  | No |

#### Triple

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| o | [BindingAttribute](#bindingattribute) |  | No |
| p | [BindingAttribute](#bindingattribute) |  | No |
| s | [BindingAttribute](#bindingattribute) |  | No |

## Database

### Creating a Virtuoso database

This can be done locally and on a server.

Pull the latest docker image to your system by running the command `docker pull openlink/virtuoso-opensource-7` in the command line.
(Optionally) check to see if you're running the latest version by using the command `docker run openlink/virtuoso-opensource-7 version`.

Create folder where you need it by using `Mkdir knox_virtuoso_db` and navigate to it by `Cd knox_virtuoso_db`.

Run the docker container specifying a password for admin access and which ports to be used:  
`docker run --name knox_virtuoso_db --interactive --tty --env DBA_PASSWORD=<> --publish 1111:1111 --publish  8890:8890 --volume pwd:/database openlink virtuoso-opensource-7:latest`

Opening `http://localhost:8890` displays the Virtuoso Conducter.
To login as admin, use username and password  to authenticate.

### Inserting Ontologies Into Virtuoso

Tunnel into the host server and cd into the volume directory /home/student.aau.dk/database.

Use Curl to install the ontology. This requires a download link for the ontology.
Use a file-transfer system to create a download link. Example: [WeTransfer](https://wetransfer.com/)

The volume will build the ontology file in the /database folder in the Virtuoso image.

Use Virtuoso's conductor interface to bulk load the ontology file. This will be done by using the Interactive SQL.

First use `ld_dir ('source-filename-or-directory', 'file name pattern', 'graph iri');`
to load the file from the specific directory.

Execute the bulk load by using `rdf_loader_run();`.

For verification use `SELECT * FROM DB.DBA.LOAD_LIST;` This can be used to check the list of data sets registered for loading, and the graph IRIs into which they will be or have been loaded.

Remember to end with a `checkpoint;`. This command MUST be run to commit the bulk loaded data to the Virtuoso database file.

Full example:![virtuoso_ontology_example.jpg](/virtuoso_ontology_example.jpg)

### Set Up Virtuoso Authorization

To ensure that only the admin or autherized users can interact with the endpoint, it is necessary to configure the autherization on Virtuoso. This exmple will use the admin user, with Username and password to authenticate.

Start by configuring the settings in Virtuoso. The first thing to do is to ensure that the `nobody` user does not have any permissions. This user is the default for unauthorized access.
To change the permissions, open the Interactive SQL in the conductor and execute this function:
`DB.DBA.RDF_DEFAULT_USER_PERMS_SET ('nobody', 0);`'.
The integer is the permission bit and follows this structure:
![permissions-virtuoso-users.png](/permissions-virtuoso-users.png)

The `dba` user will always have the proper permissions, but if another user is needed make sure to set proper permission bits by following the structure above. For example, if a user would need full permission, you would set the bit as `15(1+2+4+8)`.

To configure an endpoint to require authorization, go to the `Web Application Server` tab in the conductor. Then choose `Virtual Domains & Directories`. In there, open the folder with the correct port and find the endpoint you want authorization. The `/sparql-auth/` endpoint uses digest auth by default, but this example wants basic auth.

To configure this, we edit the `/sparql` endpoint. In here, scroll down to `Authentication Options` and:

- select `basic` in Method
- select `SPARQL Endpoint` in Realm
- Select `DB.DBA.HP_AUTH_SQL_USER` in Authentication Function

Remember to save changes.

Now when sending requests to the `/sparql/` endpoint, the user will need to set the username and password in the header of a request using basic auth.

## Watchtower

In order to pull the docker image of the repository, it is required to create an SSH key by using `ssh-keygen -t ed25519 -C "<GITHUB_MAIL>"` with your GitHub mail and add it as a deploy key in the repository. Additional information about SSH keys can be found [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent?platform=windows).  

The user who created the SSH key can then run the following command on the server to get Watchtower to run on the server. The interval indicates that Watchtower will check for updates every 30 seconds.
`docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --interval 30`

To deploy the docker image of the repository to the server a GitHub Action needs to be created. The GitHub workflow can be found as `Docker deploy image` under the Actions of this repository. The GitHub Action will trigger on pull requests to the main branch, and a docker image of the repository will be made and uploaded to the GitHub Containter Registry (ghcr.io). The image of the repository is built from the `Dockerfile` located in the repository.

## Authors

- Casper Bruun Christensen <caschr21@student.aau.dk>
- Emily Treadwell Pedersen <emiped21@student.aau.dk>
- Malthe Reipurth <mreipu21@student.aau.dk>
- Matthias Munch Jakobsen <mattja21@student.aau.dk>
- Moritz Marcus Hönscheidt <mhoens21@student.aau.dk>
- Rasmus Louie Jensen <rjen20@student.aau.dk>
