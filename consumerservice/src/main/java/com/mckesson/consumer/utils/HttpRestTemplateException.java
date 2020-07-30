/**
 * 
 */
package com.mckesson.consumer.utils;

import java.io.IOException;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import com.mckesson.consumer.service.ConsumerService;

/**
 * @author jakeermahamad
 *
 */
public class HttpRestTemplateException extends IOException {
	/**
	 * 
	 */
	private static final long serialVersionUID = 1L;

	private final static Logger log = LoggerFactory.getLogger(HttpRestTemplateException.class);
	
	private String message;
    private String statusCode;

    public HttpRestTemplateException(String statusCode, String message) {
        super(message);
        this.message = message;
        this.statusCode = statusCode;
    }

    public GeneralResponse getResponse() {
        return new GeneralResponse(statusCode, message, null);
    }

}
