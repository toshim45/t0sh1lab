package models;

import javax.persistence.*;
import java.sql.Timestamp;

/**
 * User: artikow
 * at 10/17/16 5:07 AM
 */

@Entity
public class Wallet {
    @Id
    @GeneratedValue(strategy = GenerationType.SEQUENCE)
    public Long id;

    public Double amount;

    @Version
    public Timestamp updated;
}
