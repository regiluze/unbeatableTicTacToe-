FROM scratch

COPY tic_tac_toe_game /
ENTRYPOINT ["/tic_tac_toe_game"]
