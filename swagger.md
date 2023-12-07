---
title: Database Layer Server API. v
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="">Database Layer Server API. v</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

REST API for the KNOX database.

Base URLs:

* <a href="http://http://knox-proxy01.srv.aau.dk/knox-api/">http://http://knox-proxy01.srv.aau.dk/knox-api/</a>

<h1 id="-default">Default</h1>

## getTriples

<a id="opIdgetTriples"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database \
  -H 'Accept: application/json'

```

```http
GET http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database HTTP/1.1
Host: http
Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json'
};

fetch('http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json'
}

result = RestClient.get 'http://http:/knox-proxy01.srv.aau.dk/knox-api/triples',
  params: {
  'g' => 'string'
}, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json'
}

r = requests.get('http://http:/knox-proxy01.srv.aau.dk/knox-api/triples', params={
  'g': 'http://knox_database'
}, headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://http:/knox-proxy01.srv.aau.dk/knox-api/triples', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://http:/knox-proxy01.srv.aau.dk/knox-api/triples", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /triples`

*Query graph applying filters.*

To query the whole graph, leave parameters empty.
Example: /triples?g=http://knox_database&s=subjekt1&s=subjekt2&o=object1&p=predicate1

<h3 id="gettriples-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|g|query|string|true|Target graph of the query. Currently (http://knox_ontology, http://knox_database) are valid graphs, but this could change in the future. Error responses will always contain the valid graphs, so you can just send an empty request to get the valid graphs.|
|s|query|array[string]|false|Subjects|
|o|query|array[string]|false|Objects|
|p|query|array[string]|false|Predicates|

> Example responses

> 200 Response

```json
{
  "query": "string",
  "triples": [
    {
      "o": {
        "Type": "string",
        "Value": "string"
      },
      "p": {
        "Type": "string",
        "Value": "string"
      },
      "s": {
        "Type": "string",
        "Value": "string"
      }
    }
  ]
}
```

<h3 id="gettriples-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|filtered triples response|[Result](#schemaresult)|

<aside class="success">
This operation does not require authentication
</aside>

## UpsertTriples

<a id="opIdUpsertTriples"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json'

```

```http
POST http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database HTTP/1.1
Host: http
Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "triples": [
    [
      "string"
    ]
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json'
};

fetch('http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json'
}

result = RestClient.post 'http://http:/knox-proxy01.srv.aau.dk/knox-api/triples',
  params: {
  'g' => 'string'
}, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json'
}

r = requests.post('http://http:/knox-proxy01.srv.aau.dk/knox-api/triples', params={
  'g': 'http://knox_database'
}, headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','http://http:/knox-proxy01.srv.aau.dk/knox-api/triples', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://http:/knox-proxy01.srv.aau.dk/knox-api/triples?g=http%3A%2F%2Fknox_database");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://http:/knox-proxy01.srv.aau.dk/knox-api/triples", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /triples`

*This endpoint upserts triples.*

If a new predicate is sent with an existing subject, will the existing subject be updated with the new predicate.

> Body parameter

```json
{
  "triples": [
    [
      "string"
    ]
  ]
}
```

<h3 id="upserttriples-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|g|query|string|true|Target graph of the query. Only valid graphs will be accepted. If the graph does not exist, the error response will contain the valid graphs.|
|body|body|[PostBody](#schemapostbody)|true|Triples to upsert. Each triple must contain a subject, predicate and object, in that order.|

> Example responses

> 200 Response

```json
{
  "query": "string",
  "triples": [
    {
      "o": {
        "Type": "string",
        "Value": "string"
      },
      "p": {
        "Type": "string",
        "Value": "string"
      },
      "s": {
        "Type": "string",
        "Value": "string"
      }
    }
  ]
}
```

<h3 id="upserttriples-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|response with produced query and null value for triples|[Result](#schemaresult)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_BindingAttribute">BindingAttribute</h2>
<!-- backwards compatibility -->
<a id="schemabindingattribute"></a>
<a id="schema_BindingAttribute"></a>
<a id="tocSbindingattribute"></a>
<a id="tocsbindingattribute"></a>

```json
{
  "Type": "string",
  "Value": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Type|string|false|none|none|
|Value|string|false|none|none|

<h2 id="tocS_GetTriple">GetTriple</h2>
<!-- backwards compatibility -->
<a id="schemagettriple"></a>
<a id="schema_GetTriple"></a>
<a id="tocSgettriple"></a>
<a id="tocsgettriple"></a>

```json
{
  "o": {
    "Type": "string",
    "Value": "string"
  },
  "p": {
    "Type": "string",
    "Value": "string"
  },
  "s": {
    "Type": "string",
    "Value": "string"
  }
}

```

GetTriple requires the json tags to match with the queries that are used to retrieve it.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|o|[BindingAttribute](#schemabindingattribute)|false|none|none|
|p|[BindingAttribute](#schemabindingattribute)|false|none|none|
|s|[BindingAttribute](#schemabindingattribute)|false|none|none|

<h2 id="tocS_PostBody">PostBody</h2>
<!-- backwards compatibility -->
<a id="schemapostbody"></a>
<a id="schema_PostBody"></a>
<a id="tocSpostbody"></a>
<a id="tocspostbody"></a>

```json
{
  "triples": [
    [
      "string"
    ]
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|triples|[array]|true|none|Triples is an array of triples.<br>Each triple's first element is the subject, second is the predicate and third is the object.<br>Only accepts exactly 3 elements per triple.|

<h2 id="tocS_Result">Result</h2>
<!-- backwards compatibility -->
<a id="schemaresult"></a>
<a id="schema_Result"></a>
<a id="tocSresult"></a>
<a id="tocsresult"></a>

```json
{
  "query": "string",
  "triples": [
    {
      "o": {
        "Type": "string",
        "Value": "string"
      },
      "p": {
        "Type": "string",
        "Value": "string"
      },
      "s": {
        "Type": "string",
        "Value": "string"
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|query|string|false|none|none|
|triples|[[GetTriple](#schemagettriple)]|false|none|none|

