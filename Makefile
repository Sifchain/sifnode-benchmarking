build:
	@echo
	@echo "==> Building <=="
	@ ./scripts/build.sh

run:
	@echo
	@echo "==> Running <=="
	@ ./scripts/run.sh

test:
	@echo
	@echo "==> Testing <=="
	@ rm -rf ./tmp_data
	@ mkdir ./tmp_data
	@ ./scripts/test.sh
