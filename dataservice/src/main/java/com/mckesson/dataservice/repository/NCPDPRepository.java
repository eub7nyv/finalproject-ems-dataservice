/**
 * 
 */
package com.mckesson.dataservice.repository;

import org.springframework.data.repository.CrudRepository;
import org.springframework.stereotype.Repository;

import com.mckesson.dataservice.bean.NCPDPBean;

/**
 * @author jakeermahamad
 *
 */
@Repository
public interface NCPDPRepository extends  CrudRepository<NCPDPBean, Long>{

}
