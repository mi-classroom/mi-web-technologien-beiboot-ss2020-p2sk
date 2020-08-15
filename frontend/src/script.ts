
const backendUri = 'http://localhost:8080/rest/v1/collections'

let collectionRequest = new Request(backendUri)

let handleJsonData = function(data) {
    let outputElement = document.getElementById('json-output')

    outputElement.append(JSON.stringify(data))
    /* Handle the Json data */
    console.log(data)
}

fetch(collectionRequest) 
.then(res => res.json())
.then(handleJsonData)
.catch(error => {
    console.error('Error:', error)
})