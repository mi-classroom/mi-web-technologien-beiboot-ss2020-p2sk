class BackendClient {
    private baseUri: string = 'http://localhost:8080/rest/v1/'

    private pingRoute: string = 'ping'

    private collectionsRoute: string = 'collections'

    private countQuery: string = 'count=%d'

    isOnline(handler:Function) {

    }

    request(handleData:any, handleError:any) {
        let request = new Request(this.baseUri + this.collectionsRoute)

        fetch(request)
        .then(res => res.json())
        .then(handleData)
        .catch(handleError)
    }
}

//let collectionRequest = new Request(backendUri)

let handleJsonData = function(data) {
    let outputElement = document.getElementById('json-output')

    outputElement.append(JSON.stringify(data))
    /* Handle the Json data */
    console.log(data)
}

let handleError = function(error) {
    console.error('Error:', error)
}

let client = new BackendClient

client.request(handleJsonData, handleError)
