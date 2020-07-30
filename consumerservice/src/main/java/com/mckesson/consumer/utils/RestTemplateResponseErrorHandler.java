/**
 * 
 */
package com.mckesson.consumer.utils;

import java.io.IOException;
import java.util.Arrays;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.client.ClientHttpResponse;
import org.springframework.stereotype.Component;
import org.springframework.web.client.ResponseErrorHandler;
import org.springframework.web.client.RestClientException;


/**
 * @author jakeermahamad
 *
 */
@Component
public class RestTemplateResponseErrorHandler  implements ResponseErrorHandler {
	
	private final static Logger log = LoggerFactory.getLogger(RestTemplateResponseErrorHandler.class);

	@Override
	public boolean hasError(ClientHttpResponse response) throws IOException {
		// TODO Auto-generated method stub
	return false;
	}

	@Override
	public void handleError(ClientHttpResponse response)
	            throws IOException {
	        HttpStatus statusCode = response.getStatusCode();
	        switch (statusCode.series()) {
	            case CLIENT_ERROR:
	                log.error("Client Error! response: status code: {}, status text: {}, header: {}, body: {}, charset: {}",
	                        statusCode.value(), response.getStatusText(), response.getHeaders());
	                throw new HttpRestTemplateException(String.valueOf(statusCode.value()), statusCode.value() + " " + response.getStatusText());
	            case SERVER_ERROR:
	                log.error("Server Error! response: status code: {}, status text: {}, header: {}, body: {}, charset: {}",
	                        statusCode.value(), response.getStatusText(), response.getHeaders());
	                throw new HttpRestTemplateException(String.valueOf(statusCode.value()), statusCode.value() + " " + response.getStatusText());
	            default:
	                throw new RestClientException("Unknown status code [" + statusCode + "]");
	        }
	    }
		
	}
	
	

 
