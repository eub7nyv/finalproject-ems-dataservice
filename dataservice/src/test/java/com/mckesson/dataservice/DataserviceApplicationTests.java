package com.mckesson.dataservice;

import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.web.client.RestTemplate;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

import java.util.ArrayList;
import java.util.List;

import com.mckesson.dataservice.bean.DRGPayerBean;
import com.mckesson.dataservice.bean.DRGPlanBean;
import com.mckesson.dataservice.controller.DataController;
import com.mckesson.dataservice.entity.Message;
import com.mckesson.dataservice.service.ConsumerDataService;

@ComponentScan("com.mckesson.consumer")
@SpringBootTest
class DataserviceApplicationTests {

	@Test
	void contextLoads() {
	}
	
	@Autowired
	DataController dataController;
	
	@Autowired
	ConsumerDataService consumerDataService;
	
	private static final Logger log = LoggerFactory.getLogger(DataserviceApplicationTests.class);
	
	@Test
	public void tesGreeting() {
		log.info("1...............Test Default end point....." +dataController.greeting());
		assertNotNull(dataController.greeting());
	}
	
	@Test
	public void tesHeartBeat() {
		log.info("2...............Test Default end point....." +dataController.heartBeat());
		assertNotNull(dataController.heartBeat());
	}

	
	public void testSaveConsumerDataSuccess() throws Exception {
		log.info("3...............testSaveConsumerData .....................");
		Message message = new Message();
		message.setAppName("drgpayer"); //valid application name
		message.setIncomingMessage("My Health check data");
		dataController.saveConsumerData(message);
		
	}

	@Test
	public void testSaveConsumerDataFailure() throws Exception {
		log.info("4...............testSaveConsumerDataFailure..........................: Failed: No Message Data");
		Message message = new Message();
		message.setAppName("drgpayer");//In-valid application name
		message.setIncomingMessage(" ");
		dataController.saveConsumerData(message);
		
	}
	
	@Test
    public void testgetAllDRGPayerData() throws Exception {
		log.info("5...........................Test testgetAllDRGPayerData .......");
		List<DRGPayerBean> listPayer =new ArrayList<DRGPayerBean>();
    	log.info("In sdie getAllDRGPayerData  Retrive Data >>>>>");
        listPayer= consumerDataService.getAllDRGPayerConsumerDataList();
        log.info("After Retrive Data >>>>>"+listPayer.size());
    }
	
	@Test
    public void testgetAllDRGPlanData() throws Exception {
		log.info("6...............................Test testgetAllDRGPlanData .......");
		List<DRGPlanBean> listPayer =new ArrayList<DRGPlanBean>();
    	log.info("In sdie getAllDRGPayerData  Retrive Data >>>>>");
        listPayer= consumerDataService.getAllDRGPlanConsumerDataList();
        log.info("After Retrive Data >>>>>"+listPayer.size());
    }
	
	
	
}
