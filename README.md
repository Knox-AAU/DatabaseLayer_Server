# DatabaseLayer_Server

Go REST API with CRUD operations for Knox database

# Table of Contents

- [DatabaseLayer\_Server](#databaselayer_server)
- [Table of Contents](#table-of-contents)
  - [Branch protection rules](#branch-protection-rules)
  - [Debug Github Actions](#debug-github-actions)
  - [Generate new documentation based on code](#generate-new-documentation-based-on-code)
  - [Servers and access](#servers-and-access)
  - [Accessing our server and API](#accessing-our-server-and-api)
    - [Access our Virtuoso database](#access-our-virtuoso-database)
    - [Access our database layer API](#access-our-database-layer-api)
  - [Deploy server](#deploy-server)
  - [Creating a Virtuoso database](#creating-a-virtuoso-database)
  - [Watchtower](#watchtower)
  - [Authors](#authors)

## Branch protection rules

The main branch of the repository is protected by branch protection rules, which ensures the code can build and is tested before it can be merged to the main branch. The branch protection rule requires pull requests before merging and status checks to pass before merging. The status checks are defined in the `test.yml` file from the Test workflow in GitHub Actions.

## Debug Github Actions

See this [CLI tool](https://github.com/nektos/act) to run Github Actions locally 

## Generate new documentation based on code

Run `swagger generate spec -m -o ./swagger.yaml` from the terminal or directly from the `main.go` file.

After generating the yaml file, run `redocly build-docs swagger.yaml` from the terminal, which will generate the updated html docs.

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

You can access the server from the command line with `ssh <STUDENT_MAIL>@<SERVER_NAME> -L <PORT>:localhost:<PORT>`.

## Accessing our server and API

### Access our Virtuoso database

The code for this project (including tests) accesses Virtuoso on port 8890, which is the same port to access it on the server.

`ssh <STUDENT_MAIL>@knox-kb01.srv.aau.dk -L 8890:localhost:8890`

### Access our database layer API

Your port is 8000 and the API on the server is on 8081.

`ssh <STUDENT_MAIL>@knox-kb01.srv.aau.dk -L 8000:localhost:8081`

## Deploy server

`docker run -p 8081:8080 --add-host=host.docker.internal:host-gateway -e VIRTUOSO_SERVER_URL=http://host.docker.internal:8890/sparql -e VIRTUOSO_GRAPH_URI=http://localhost:8890/DAV -e VIRTUOSO_TEST_GRAPH_URI=http://testing -d ghcr.io/knox-aau/databaselayer_server:main`

## Creating a Virtuoso database

This can be done locally and on a server.

Pull the latest docker image to your system by running the command `docker pull openlink/virtuoso-opensource-7` in the command line.
(Optionally) check to see if you're running the latest version by using the command `docker run openlink/virtuoso-opensource-7 version`.

Create folder where you need it by using `Mkdir knox_virtuoso_db` and navigate to it by `Cd knox_virtuoso_db`.

Run the docker container specifying a password for admin access and which ports to be used:  
`docker run --name knox_virtuoso_db --interactive --tty --env DBA_PASSWORD=qzu49svh --publish 1111:1111 --publish  8890:8890 --volume pwd:/database openlink virtuoso-opensource-7:latest`

Opening `http://localhost:8890` displays the Virtuoso Conducter. 
To login as admin, use the username ***dba*** and password ***qzu49svh*** which was specified in the previous command.

## Watchtower

In order to pull the docker image of the repository, it is required to create an SSH key by using `ssh-keygen -t ed25519 -C "<GITHUB_MAIL>"` with your GitHub mail and add it as a deploy key in the repository. Additional information about SSH keys can be found [here](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent?platform=windows).  

The user who created the SSH key can then run the following command on the server to get Watchtower to run on the server. The interval indicates that Watchtower will check for updates every 30 seconds. 
`docker run -d --name watchtower -v /var/run/docker.sock:/var/run/docker.sock containrrr/watchtower --interval 30`

To deploy the docker image of the repository to the server a GitHub Action needs to be created. The GitHub workflow can be found as `Docker deploy image` under the Actions of this repository. The GitHub Action will trigger on pull requests to the main branch, and a docker image of the repository will be made and uploaded to the GitHub Containter Registry (ghcr.io). The image of the repository is built from the `Dockerfile` located in the repository.   

## Authors

Casper Bruun Christensen (caschr21@student.aau.dk)  
Emily Treadwell Pedersen (emiped21@student.aau.dk)  
Malthe Reipurth (mreipu21@student.aau.dk)  
Matthias Munch Jakobsen (mattja21@student.aau.dk)  
Moritz Marcus Hönscheidt (mhoens21@student.aau.dk)  
Rasmus Louie Jensen (rjen20@student.aau.dk)


