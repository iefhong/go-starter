FROM golang:1.14 AS development

# https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md#walk-through
ENV GOBIN /app/bin
ENV PATH $GOBIN:$PATH

# Add the official postgres repo to install the matching postgresql-client tools of your stack
# see https://wiki.postgresql.org/wiki/Apt
# run lsb_release -c inside the container to pick the proper repository flavor
# e.g. stretch=>stretch-pgdg, buster=>buster-pgdg
RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ buster-pgdg main" \
    | tee /etc/apt/sources.list.d/pgdg.list \
    && wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc \
    | apt-key add -

# Install required system dependencies
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    postgresql-client-12 \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Install the same version of pg_formatter as used in your editors, as of 2020-03 thats v4.2
# https://github.com/darold/pgFormatter/releases
# https://github.com/bradymholt/vscode-pgFormatter/commits/master
RUN wget https://github.com/darold/pgFormatter/archive/v4.2.tar.gz \
    && tar xzf v4.2.tar.gz \
    && cd pgFormatter-4.2 \
    && perl Makefile.PL \
    && make && make install

# Yes no maybe. This is strange. Although all default shells are bash and bash has been set as the shell for yarn/npm to use, 
# it still runs everything as /bin/sh for some weird reason. Let's make sure it doesn't. Naughty yarn. 
# RUN rm /bin/sh \ 
#     && ln -s /bin/bash /bin/sh

# We use ssmtp (provides a sendmail binary) to relay emails to servers (must be custom configured)
# https://blog.philippklaus.de/2011/03/set-up-sending-emails-on-a-local-system-by-transfering-it-to-a-smtp-relay-server-smarthost
# Configure by mounting /etc/ssmtp/ssmtp.conf
# RUN apt-get update \
#     && apt-get install -y ssmtp \
#     && apt-get clean \
#     && rm -rf /var/lib/apt/lists/*

# Install required system dependencies
# E.g.
# RUN set -e \
#     && apt-get update \
#     && apt-get install -y --no-install-recommends \
#     imagemagick \
#     && rm -rf /var/lib/apt/lists/*

# Comment int, if using psql cli in tests
# TESTS_ONLY: We need an psql client (any version, fuckit) to execute tests that utilize the psql cli
# RUN set -e \
#     && apt-get update \
#     && apt-get install -y --no-install-recommends \
#     postgresql-client \
#     && rm -rf /var/lib/apt/lists/*