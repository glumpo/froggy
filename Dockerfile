## ----------------------------------------------------------------------------
## Build
## ----------------------------------------------------------------------------

FROM golang:1.19.1-bullseye AS build

WORKDIR /app
COPY . .

ARG version
ENV VERSION ${version:-'0.0.0-develop'}
RUN make froggy

## ----------------------------------------------------------------------------
## Deploy
## ----------------------------------------------------------------------------

FROM debian:11.5

WORKDIR /opt/froggy
COPY --from=build /app/dist/* .
COPY --from=build /app/etc/config.toml ./config.toml

RUN useradd nonroot --user-group --no-create-home
RUN chown -R nonroot:nonroot /opt/froggy && chmod -R 755 /opt/froggy

USER nonroot:nonroot

ENTRYPOINT ["/opt/froggy/froggy"]

