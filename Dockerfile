FROM ubuntu:14.04

COPY tic_tac_toe_game /
ENTRYPOINT ["/tic_tac_toe_game"]
