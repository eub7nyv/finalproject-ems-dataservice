/**
 * 
 */
package com.mckesson.dataservice.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author jakeermahamad
 *
 */

@RestController
public class DataController {
	
	private static final Logger log = LoggerFactory.getLogger(DataController.class);
	
	@GetMapping("/greeting")
    public String greeting() {
        return "Hello From  DataController";
    }

}
