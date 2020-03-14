'use strict'
function login(){
	email = document.getElementById('login-email').value;
	password = document.getElementById('login-password').value;
	console.log(email);
	console.log(password);

	var requestBody = JSON.stringify({
	    Email: email,
	    CreatePassword: password
	});

	fetch({{.Config.LoginAttemptedPageURL}}, {
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		if (!data.Success) {
			handleLoginMessage(data.Message)
			return 
		}
		window.location.replace("/terminal");
	}).catch(err => {
		console.log(err)
	});
}

function handleLoginMessage(err) {
	loginErrorBox = document.getElementById('loginErrorMessage');
	loginErrorBox.innerHTML = err;
}

function signup() {
	firstName = document.getElementById("first_name").value;
	lastName = document.getElementById("last_name").value;
	createPassword = document.getElementById("create-password").value;
	confirmPassword = document.getElementById("confirm-password").value;
	email = document.getElementById("email").value;

	console.log(firstName)
	console.log(lastName)
	console.log(createPassword)
	console.log(confirmPassword)
	console.log(email)

	if (confirmPassword == ""){
		handleSignUpError("confirm password not filled in")	
		return
	} else if (createPassword == "") {
		handleSignUpError("create password not filled in")	
		return
	} else if (firstName == "") {
		handleSignUpError("first name not filled in")	
		return
	} else if (lastName == "") {
		handleSignUpError("last name not filled in")	
		return
	} else if (email == "") {
		handleSignUpError("Email not filled in")	
		return
	} else if (createPassword != confirmPassword) {
		handleSignUpError("Passwords must match")	
		return
	}

	var requestBody = JSON.stringify({
		FirstName: firstName,
		LastName: lastName,
		CreatePassword: createPassword,
		ConfirmPassword: confirmPassword,
		Email: email
	});

	console.log(requestBody);

	fetch({{.Config.CreateAccountURL}}, {
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		if (!data.Success) {
			handleSignUpError(data.Message)
			return
		}

		window.location.replace("/terminal");
	}).catch(err => {
		console.log(err)
	});
}

function handleSignUpError(err){
	signUpErrorBox = document.getElementById('signUpErrorMessage');
	signUpErrorBox.innerHTML = err;
}


