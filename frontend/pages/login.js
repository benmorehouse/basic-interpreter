"use strict";

function login(){
	email = document.getElementById('login-email').value;
	password = document.getElementById('login-password').value;
	console.log(email);
	console.log(password);

	var requestBody = JSON.stringify({
	    Email: email,
	    Password: password
	});
	
	fetch({{.Config.LoginURL}}, {
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		console.log(data);
		window.location.replace("/terminal");
	}).catch(err => {
		console.log(err)
	});
}

function signup(){
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
		errorMessage("confirm password not filled in")	
		return
	} else if (createPassword == "") {
		errorMessage("first name not filled in")	
		return
	} else if (firstName == "") {
		errorMessage("first name not filled in")	
		return
	} else if (lastName == "") {
		errorMessage("Last name not filled in")	
		return
	} else if (email == "") {
		errorMessage("Email not filled in")	
		return
	} else if (createPassword != confirmPassword) {
		errorMessage("Passwords must match")	
		return
	}

	var requestBody = JSON.stringify({
		FirstName: firstName,
		LastName: lastName,
		CreatePassword: createPassword,
		ConfirmPassword: confirmPassword,
		Email: email
	});

	fetch({{.Config.CreateAccountURL}}, {
		method:"POST",
		credentials:"include",
		body: requestBody,		
	}).then(res => {
		return res.json();
	}).then(data => {
		console.log(data)
		window.location.replace("/terminal");
	}).catch(err => {
		console.log(err)
	});
}

function errorMessage(err){
	console.log(err)
}


