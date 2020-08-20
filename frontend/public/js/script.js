var BackendClient = /** @class */ (function () {
    function BackendClient() {
        this.baseUri = 'http://localhost:8080/rest/v1/';
        this.pingRoute = 'ping';
        this.collectionsRoute = 'collections';
        this.countQuery = 'count=%d';
    }
    BackendClient.prototype.isOnline = function (handler) {
    };
    BackendClient.prototype.request = function (handleData, handleError) {
        var request = new Request(this.baseUri + this.collectionsRoute);
        fetch(request)
            .then(function (res) { return res.json(); })
            .then(handleData)["catch"](handleError);
    };
    return BackendClient;
}());
//let collectionRequest = new Request(backendUri)
var handleJsonData = function (data) {
    var outputElement = document.getElementById('json-output');
    outputElement.append(JSON.stringify(data));
    /* Handle the Json data */
    console.log(data);
};
var handleError = function (error) {
    console.error('Error:', error);
};
var client = new BackendClient;
client.request(handleJsonData, handleError);
