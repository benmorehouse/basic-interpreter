
function login(){
	username = document.getElementById('login-password').value;
	password = document.getElementById('login-password').value;
	console.log(username);
	console.log(password);

	var requestBody = JSON.stringify({
	    Username: username,
	    Password: password,
	});
	
	var fetchURL = "/" + {{.Config.LoginURL}};
	console.log("fetching the following:", fetchURL)
	
	fetch(fetchURL, {
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		console.log(data)
	}).catch(err => {
		console.log(err)
	});
}

