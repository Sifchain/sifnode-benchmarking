build:
	@echo
	@echo "==> Building <=="
	@./scripts/build.sh

run:
	@echo
	@echo "==> Running <=="
	@./scripts/run.sh

test:
	@echo
	@echo "==> Testing <=="
	@./scripts/test.sh

test-save-artifacts:
	@echo
	@echo "==> Testing <=="
	@./scripts/test_save_artifacts.sh

clean:
	@echo
	@echo "==> Cleaning <=="
	@./scripts/clean.sh
