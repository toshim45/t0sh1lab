package com.talentmgr.api.domain.profile;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import reactor.core.publisher.Flux;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

/**
 * Created by artikow on 9/30/17.
 */
@RestController
@RequestMapping("/profiles")
public class ProfileController {
	@RequestMapping(method= RequestMethod.GET)
	public Flux<Profile> findAll(){
		ArrayList<Profile> profiles = new ArrayList<>();
		profiles.add(new Profile("me-1", new Date()));
		profiles.add(new Profile("me-2", new Date()));
		return Flux.fromIterable(profiles);
	}
}
