/**
 * 
 */
package com.mckesson.dataservice.controller;

import java.util.ArrayList;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.mckesson.dataservice.bean.DRGPayerBean;
import com.mckesson.dataservice.bean.DRGPlanBean;
import com.mckesson.dataservice.bean.NCPDPBean;
import com.mckesson.dataservice.entity.ConsumerConstants;
import com.mckesson.dataservice.entity.Message;
import com.mckesson.dataservice.service.ConsumerDataService;

/**
 * @author jakeermahamad
 *
 */

@RestController
public class DataController {
	
	private static final Logger log = LoggerFactory.getLogger(DataController.class);
	
	
	 @Autowired
	 ConsumerDataService consumerDataService;
	
	@GetMapping("/greeting")
    public String greeting() {
        return "Hello From  DataController";
    }
	
	@GetMapping("/mckesson")
    public String heartBeat() {
        return "Application Health -OK";
    }
	
	 @PostMapping(path = "/mckesson/saveConsumerData")
	    public String saveConsumerData(@RequestBody Message msg) throws Exception {
	         String topic =msg.getAppName();
	    	 log.info("Topic Name>>>>>"+topic+"<< In Coming Message >>>"+msg.getIncomingMessage());
	    	 String statusMessage = null;
	        switch (topic) {
	        case ConsumerConstants.TOPIC_DRG_PAYER:
	        	statusMessage = consumerDataService.insertDRGPayerConsumerData(msg.getIncomingMessage()); 
	            break;
	        case ConsumerConstants.TOPIC_DRG_PLAN:
	        	statusMessage = consumerDataService.insertDRGPlanConsumerData(msg.getIncomingMessage());
	            break;
	        case ConsumerConstants.TOPIC_NCPDP:
	        	statusMessage = consumerDataService.insertDRGPlanConsumerData(msg.getIncomingMessage());
	            break;
	        default:
	        	statusMessage = topic+": Topic Name is not configured";
	            break;
	    }
	        
	        
	        
	     
	        return statusMessage;


	    }

	    
	    
	    @RequestMapping(value= "/mckesson/listofMessages/drgpayer")
	    public List<?> getAllDRGPayerData() throws Exception {
	    	List<DRGPayerBean> listPayer =new ArrayList<DRGPayerBean>();
	    	log.info("In sdie getAllDRGPayerData  Retrive Data >>>>>");
	        listPayer= consumerDataService.getAllDRGPayerConsumerDataList();
	        log.info("After Retrive Data >>>>>");
	    	
	        return listPayer;
	    }
	    
	    @RequestMapping(value= "/mckesson/listofMessages/drgplan")
	    public List<DRGPlanBean> getAllDrgPlanData() throws Exception {
	    	log.info("In sdie getAlldrgPlanData  Retrive Data >>>>>");
	    	List<DRGPlanBean> listPayer =new ArrayList<DRGPlanBean>();
	    	listPayer=  consumerDataService.getAllDRGPlanConsumerDataList();
	        log.info("After Retrive Data >>>>>");
	        return listPayer;
	    }
	    
	    @RequestMapping(value= "/mckesson/listofMessages/ncpdp")
	    public List<NCPDPBean> getAllNcpdpPlanData() throws Exception {
	    	log.info("In sdie getAllNcpdpPlanData  Retrive Data >>>>>");
	    	List<NCPDPBean> listPayer =new ArrayList<NCPDPBean>();
	    	listPayer=  consumerDataService.getAllNCPDPConsumerDataList();
	        log.info("After Retrive Data >>>>>");
	        return listPayer;
	    }


}
