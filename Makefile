PREFIX=/usr/local
DESTDIR=
GOFLAGS=
BINDIR=${PREFIX}/bin

NAMESERVER_SRCS = $(wildcard apps/nameserver/*.go nameserver/*.go)
BROKER_SRCS = $(wildcard apps/broker/*.go broker/*.go)
PRODUCER_SRCS = $(wildcard apps/client/producer/*.go client/producer/*.go) 
CONSUMER_SRCS = $(wildcard apps/client/consumer/*.go client/consumer/*.go) 

APPS = nameserver broker 
CLIENT_APPS = producer consumer
BLDDIR = build

all: $(APPS) $(CLIENT_APPS)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${GOFLAGS} -o $(abspath $@) ./$*

$(APPS): %: $(BLDDIR)/apps/%
$(CLIENT_APPS): %: $(BLDDIR)/apps/client/%

$(BLDDIR)/apps/nameserver: $(NAMESERVER_SRCS)
$(BLDDIR)/apps/broker: $(BROKER_SRCS)
$(BLDDIR)/apps/client/producer: $(PRODUCER_SRCS)
$(BLDDIR)/apps/client/consumer: $(CONSUMER_SRCS)

clean:
	rm -fr $(BLDDIR)

.PHONY: install clean all
.PHONY: $(APPS)

install: $(APPS) $(CLIENT_APPS)
	install -m 755 -d ${DESTDIR}${BINDIR}
	install -m 755 $(BLDDIR)/apps/nameserver ${DESTDIR}${BINDIR}/nameserver
	install -m 755 $(BLDDIR)/apps/broker ${DESTDIR}${BINDIR}/broker
	install -m 755 $(BLDDIR)/apps/client/producer ${DESTDIR}${BINDIR}/producer
	install -m 755 $(BLDDIR)/apps/client/consumer ${DESTDIR}${BINDIR}/consumer
