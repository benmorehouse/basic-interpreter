# Filestore

this is documentation on a specific table within the database.

Filestore: 
		id varchar(30),
		userid varchar(30),
		file LONGBLOB,

---

filestore is simply a way to hold files for each respective user.

the id corresponds to the path the user has created to get to the file, as well
as the filename. 
The userid is the user's id.
File is a longblob, which can hold up to 2 gigabytes of information. Plenty for a basic file. 

