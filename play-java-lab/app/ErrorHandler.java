import play.Configuration;
import play.Environment;
import play.api.OptionalSourceMapper;
import play.api.UsefulException;
import play.api.routing.Router;
import play.http.DefaultHttpErrorHandler;
import play.http.HttpErrorHandler;
import play.libs.Json;
import play.mvc.Http;
import play.mvc.Result;
import play.mvc.Results;

import javax.inject.Inject;
import javax.inject.Provider;
import javax.inject.Singleton;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.CompletionStage;

/**
 * Created by awibowo on 18/10/2016.
 */
@Singleton
public class ErrorHandler implements HttpErrorHandler {

    @Override
    public CompletionStage<Result> onClientError(Http.RequestHeader requestHeader, int code, String message) {
        return CompletableFuture.completedFuture(
                Results.status(code, "A client error occurred: " + message)
        );
    }

    @Override
    public CompletionStage<Result> onServerError(Http.RequestHeader requestHeader, Throwable throwable) {
        return CompletableFuture.completedFuture(
                Results.internalServerError(Json.toJson(new ErrorResponseMessage(throwable.getMessage())))
        );
    }
}

class ErrorResponseMessage {
    public String message;

    public ErrorResponseMessage(String message) {
        this.message = message;
    }
}