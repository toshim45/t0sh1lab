import play.filters.cors.CORSFilter;
import play.http.DefaultHttpFilters;

import javax.inject.Inject;
import javax.inject.Singleton;

/**
 * This class configures filters that run on every request. This
 * class is queried by Play to get a list of filters.
 *
 * Play will automatically use filters from any class called
 * <code>Filters</code> that is placed the root package. You can load filters
 * from a different class by adding a `play.http.filters` setting to
 * the <code>application.conf</code> configuration file.
 */
@Singleton
public class Filters extends DefaultHttpFilters {

    @Inject
    public Filters(CORSFilter corsFilter) {
        super(corsFilter);
    }
}
