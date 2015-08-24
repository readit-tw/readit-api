Readit API [![Build Status](https://travis-ci.org/readit-tw/readit-api.svg?branch=master)](https://travis-ci.org/readit-tw/readit-api)
==========

API for Readit

Setup Reference links
	-- -- You can install all dependencies by running "go get -v ./..." from the root of your workspace.
		-- "go get -v ./..."
	-- Install eclipse goclipse plugin
	-- https://github.com/golang/go/wiki/GithubCodeLayout
	-- https://github.com/GoClipse/goclipse/blob/latest/documentation/UserGuide.md 
-
	

	- client : ssh-keygen on ~/.ssh wil create public key in id_rsa.pub
	- servers :  copy this public key in   ~/.ssh folder authorized_keys file
	- client : create config file in  ~/.ssh file with content as 
	      Host <alias>
	     	 HostName <host name>
	     	 UserName <user name>
	     	 
	- client : cross compiler go install
	- client : Makefile is there then do make deploy 
	
	-------For info if server get restart then automatically start our process
	     file : /etc/init/readit-api 
	    content : exec /path/to/binary/file
	    
MOngo:
	-- connect to server and start 'mongo'
	-- to remove all documents 
		-- db.resources.remove({}) 