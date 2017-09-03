package controllers;

import models.Wallet;
import play.db.jpa.Transactional;
import play.libs.Json;
import play.mvc.Result;
import play.mvc.Controller;
import play.db.jpa.JPAApi;

import javax.inject.Inject;
import javax.persistence.TypedQuery;
import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Root;

/**
 * User: artikow
 * at 10/16/16 10:50 AM
 */
public class WalletController extends Controller {
    @Inject
    private JPAApi jpaApi;

    @Transactional(readOnly = true)
    public Result index(Double balance) {
        CriteriaBuilder cb = jpaApi.em().getCriteriaBuilder();
        CriteriaQuery<Wallet> cq = cb.createQuery(Wallet.class);
        Root<Wallet> root = cq.from(Wallet.class);
        CriteriaQuery<Wallet> all = cq.select(root);
        if (balance != null) cq.where(cb.equal(root.get("balance"), balance));
        TypedQuery<Wallet> allQuery = jpaApi.em().createQuery(all);
        return ok(Json.toJson(allQuery.getResultList()));
    }

    @Transactional(readOnly = true)
    public Result detail(Long id) {
        Wallet wallet = jpaApi.em().find(Wallet.class, id);
        if (wallet == null) return notFound();
        return ok(Json.toJson(wallet));
    }
}
