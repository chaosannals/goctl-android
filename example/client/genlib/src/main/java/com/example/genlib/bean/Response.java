package com.example.genlib.bean;


import com.google.gson.annotations.Expose;
import com.google.gson.annotations.SerializedName;

public class Response {
	
	@SerializedName("message")
	private String message; 
	
	public String getMessage() {
		return message;
	}

	public void setMessage(String message) {
		this.message = message;
	}
	
}
