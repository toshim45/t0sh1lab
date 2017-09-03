package models;

import java.util.ArrayList;
import java.util.List;

/**
 * User: artikow
 * at 9/20/16 12:03 AM
 */
public class ScoreStrategicInitiative {
    public long id;
    public double processScore;
    public double impactScore;
    public double totalScore;

    public ScoreStrategicInitiative(long id, double processScore, double impactScore, double totalScore) {
        this.id = id;
        this.processScore = processScore;
        this.impactScore = impactScore;
        this.totalScore = totalScore;
    }

    public static List<ScoreStrategicInitiative> findAll(Long periodIdx, String type) {
        ArrayList<ScoreStrategicInitiative> scores = new ArrayList<>();
        scores.add(new ScoreStrategicInitiative(periodIdx, 93.211d, 38.128d, 58.912d));
        return scores;
    }
}
