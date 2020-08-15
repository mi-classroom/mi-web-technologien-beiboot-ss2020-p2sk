var backendUri = 'http://localhost:8080/rest/v1/collections';
var collectionRequest = new Request(backendUri);
var handleJsonData = function (data) {
    var outputElement = document.getElementById('json-output');
    outputElement.append(JSON.stringify(data));
    /* Handle the Json data */
    console.log(data);
};
fetch(collectionRequest)
    .then(function (res) { return res.json(); })
    .then(handleJsonData)["catch"](function (error) {
    console.error('Error:', error);
});
