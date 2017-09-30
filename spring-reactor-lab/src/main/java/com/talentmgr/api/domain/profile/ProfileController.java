package com.talentmgr.api.domain.profile;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by artikow on 9/30/17.
 */
@RestController
@RequestMapping("/profiles")
public class ProfileController {
	@RequestMapping(method= RequestMethod.GET)
	public List<Profile> findAll(){
		return new ArrayList<Profile>();
	}
}
