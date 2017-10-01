package com.talentmgr.api.domain.profile;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.HttpClientErrorException;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;
import reactor.core.publisher.MonoSink;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import java.util.UUID;

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

	@RequestMapping(method = RequestMethod.GET, value = "/{id}")
	public Mono<Profile> findOne(@PathVariable("id") UUID id){
		return Mono.just(new Profile("me-3", new Date()));
	}

	@RequestMapping(method = RequestMethod.POST)
	@ResponseStatus(value = HttpStatus.CREATED)
	public Mono<Boolean> createOne(@RequestBody Profile profile) {
		return Mono.empty();
	}
}
