function Signup(userField, passField) {
	let headers = {
		'Content-Type': 'application/json',
		username: userField,
		password: passField,
	};

	axios
		.get('http://localhost:3000/register', {}, [headers])
		.then((response) => {
			// Handle success
			console.log('Response:', response.data);
			alert('Signup successful!');
		})
		.catch((error) => {
			// Handle error
			console.error('Error:', error);
			alert('Signup failed.');
		});
}
