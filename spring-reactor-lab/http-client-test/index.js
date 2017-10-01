const httpClient = require('axios');
const uuidv4 = require('uuid/v4');

id = uuidv4();

httpClient.get('http://localhost:8080/profiles/'+id)
	.then(function (response) {
  		if (response.status != 200) {
			throw new Error('http status: ' + response.status)
		}
		console.log(response.data);
  	});


httpClient.get('http://localhost:8080/profiles')
	.then(function (response) {
		if (response.status != 200) {
			throw new Error('http status: ' + response.status)
		}
		console.log(response.data);
	});

httpClient.post('http://localhost:8080/profiles', {
	description: "me-req-1"
}).then(function (response) {
	if (response.status != 201) {
			throw new Error('http status: ' + response.status)
		}
});
