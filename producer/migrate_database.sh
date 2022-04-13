#!/usr/bin/env bash

readonly SCRIPT_ARGS_STR="$@"
readonly SCRIPT_RELATIVE_PATH="$(dirname $0)"
readonly SCRIPT_ABSOLUTE_PATH="$(realpath $SCRIPT_RELATIVE_PATH)"
readonly SCRIPT_NAME="$(basename $0)"
readonly CURRENT_FOLDER="$(basename $PWD)"
readonly RUNNING_FROM="${PWD}"
readonly POSTGRESQL_URL="postgres://bucketeer:bucketeer_pass@localhost:5432/bucketeer_db?sslmode=disable"

function up_database() {
  echo -e "up called with ${@}"
  "${SCRIPT_ABSOLUTE_PATH}/../../migrate" -database ${POSTGRESQL_URL} -path "${SCRIPT_ABSOLUTE_PATH}/db/migrations" -verbose up
}

function down_database() {
  echo -e "down called with ${@}"
  "${SCRIPT_ABSOLUTE_PATH}/../../migrate" -database ${POSTGRESQL_URL} -path "${SCRIPT_ABSOLUTE_PATH}/db/migrations" -verbose down
}

if [[ "${#SCRIPT_ARGS_STR}" -eq 0 ]]; then
  echo -e "\033[1;31mThis script has to be called with a subcommand name"
  echo -e "\t\033[1;31mEither 'up' or 'down'"
  exit 1
fi

case "${1}" in
  UP | Up | up)
    up_database $SCRIPT_ARGS_STR
  ;;
  DOWN | Down | down)
    down_database $SCRIPT_ARGS_STR
  ;;
esac

