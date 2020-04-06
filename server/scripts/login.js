'use strict'

var loginEmail = "";
var loginPassword = "";

var firstName = "";
var lastName = "";
var createPassword = "";
var confirmPassword = "";
var createEmail = "";

function onInputLoginEmail(e) {
	
	loginEmail = e;
}

function onInputLoginPassword(e) {

	loginPassword = e;
}

function onInputFirstname(e) {
	
	firstName = e;
}

function onInputLastname(e) {

	lastName = e;
}

function onInputPassword(e) {

	createPassword = e
	if (createPassword != confirmPassword) {
		handleSignUpError("Passwords must match")	
		return
	}

}

function onInputConfirmpassword(e) {

	confirmPassword = e;
	if (createPassword != confirmPassword) {
		handleSignUpError("Passwords must match")	
		return
	}
}

function onInputCreateEmail(e) {

	createEmail = e;
}

function login(loginEndpoint){
	
	console.log("Attempting a user login...")

	var requestBody = JSON.stringify({
	    Email: loginEmail,
	    CreatePassword: loginPassword
	});
	
	// need to figure out a way to get the endpoints needed for this.
	fetch(loginEndpoint, {  
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

function signup(signUpEndpoint) {

	console.log("Attempting a user signup...")
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
	} else if (createEmail == "") {
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
		Email: createEmail
	});

	fetch(signUpEndpoint, {
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


