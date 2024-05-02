var grpc = require('grpc');
var messages = require('./message_pb');
var services = require('./message_grpc_pb');

const { Empty } = require('google-protobuf/google/protobuf/empty_pb');

function main() {
	var client = new services.MessageServiceClient('localhost:8081', grpc.credentials.createInsecure());
	var request = new messages.MessageRequest();
	request.setLabel('test');

	client.updateMessage(request, function(err, response) {
		console.log('result-update-error:', err)
		console.log('result-update-label:', response.getLabel())
		console.log('result-update-created:', response.getCreated())
	});

	client.checkMessage(new Empty(), function(err, response) {
		console.log('result-check-error:', err)
		console.log('result-check-response:', response.getStatus())
	});

	client.listServingStatus(new Empty(), function(err, response){
		console.log('result-list-serving-status:', err)
		console.log('result-list-serving-status-response:', response.toObject())
		console.log('result-list-serving-status-response-1:', response.toObject().statusMap[1])
	});
}

main();
