package controllers;

import models.ScoreStrategicInitiative;
import play.libs.Json;
import play.mvc.Controller;
import play.mvc.Result;

/**
 * Created by awibowo on 19/09/2016.
 */
public class DashboardController extends Controller {
    public Result index(Long id, String type){
        return ok(Json.toJson(ScoreStrategicInitiative.findAll(id, type)));
    }
}