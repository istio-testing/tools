
define kubetype-gen =
$(shell go run ./main.go --input-dirs=$$(go list ${1}) --output-package="$$(go list .)/testdata/test_output/${2}" --output-base=. --trim-path-prefix=istio.io/tools/cmd/kubetype-gen -h "./boilerplate.go.txt" 2>&1)
endef

test: run-negative-tests run-positive-tests

check-env:
	@if ! git diff --quiet HEAD -- testdata/test_input; then \
	    echo "FAIL: test inputs contain uncommitted changes. Please commit changes in ./testdata/test_input."; \
	    false; \
	fi

run-positive-tests: run-positive-defaults-test run-positive-types-test

run-positive-defaults-test: check-env
	@echo "Running $@"
	$(eval $@_output = $(call kubetype-gen,./testdata/test_input/positive/defaults,defaults))
	@if echo "${$@_output}" | grep -q "Errors occurred while scanning input.  See previous output for details."; then \
	    echo "kubetype-gen command failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! git diff --quiet HEAD -- testdata/test_output || [ -n "$$(git ls-files -o -- testdata/test_output)" ]; then \
	    echo "FAIL: output files have changed"; \
	    false; \
	fi

run-positive-types-test: check-env
	@echo "Running $@"
	$(eval $@_output = $(call kubetype-gen,./testdata/test_input/positive/types,types))
	@if echo "${$@_output}" | grep -q "Errors occurred while scanning input.  See previous output for details."; then \
	    echo "kubetype-gen command failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Invalid value specified for +kubetype-gen:kubeType in type .*/test_input/positive/types.EmptyKubeType.  Using default name EmptyKubeType."; then \
	    echo "Missing warning for types.EmptyKubeType"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! git diff --quiet HEAD -- testdata/test_output || [ -n "$$(git ls-files -o -- testdata/test_output)" ]; then \
	    echo "FAIL: output files have changed"; \
	    false; \
	fi


run-negative-tests: run-negative-defaults-empty-group-test run-negative-defaults-invalid-group-version-test run-negative-types-test

run-negative-defaults-empty-group-test: check-env
	@echo "Running $@"
	$(eval $@_output = $(call kubetype-gen,./testdata/test_input/negative/defaults/emptygroup,fail))
	@if ! echo "${$@_output}" | grep -q "Errors occurred while scanning input.  See previous output for details."; then \
	    echo "kubetype-gen command should have failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Invalid Group/Version for package istio.io/tools/cmd/kubetype-gen/testdata/test_input/negative/defaults/emptygroup, Group not specified for Group/Version: groupversion"; then \
	    echo "negative/defaults/emptygroup test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi

run-negative-defaults-invalid-group-version-test: check-env
	@echo "Running $@"
	$(eval $@_output = $(call kubetype-gen,./testdata/test_input/negative/defaults/invalidgroupversion,fail))
	@if ! echo "${$@_output}" | grep -q "Errors occurred while scanning input.  See previous output for details."; then \
	    echo "kubetype-gen command should have failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Could not calculate Group/Version for package .*/test_input/negative/defaults/invalidgroupversion: invalid group version 'group/version/version' specified: unexpected GroupVersion string: group/version/version"; then \
	    echo "negative/defaults/invalidgroupversion test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi

run-negative-types-test: check-env
	@echo "Running $@"
	$(eval $@_output = $(call kubetype-gen,./testdata/test_input/negative/types,fail))
	@if ! echo "${$@_output}" | grep -q "Errors occurred while scanning input.  See previous output for details."; then \
	    echo "kubetype-gen command should have failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -qE "duplicate kube types specified for .*/test_output/fail/group/version.AGoodType.  duplicated by: \[istio.io/tools/cmd/kubetype-gen/testdata/test_input/negative/types.(AGoodType|DuplicateKubeType) istio.io/tools/cmd/kubetype-gen/testdata/test_input/negative/types.(AGoodType|DuplicateKubeType)\]"; then \
	    echo "types.DuplicateKubeType test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Invalid Group/Version for type .*/test_input/negative/types.EmptyGroup: groupversion"; then \
	    echo "types.EmptyGroup test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Could not calculate Group/Version for type .*/test_input/negative/types.InvalidGroupVersion: invalid group version 'group/version/version' specified"; then \
	    echo "types.InvalidGroupVersion test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Invalid Group/Version for type .*/test_input/negative/types.NoGroupVersion: <nil>"; then \
	    echo "types.NoGroupVersion test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -qE "Overlapping packages for Group/Versions group2(.name.io)?/version and group2(.name.io)?/version"; then \
	    echo "types.OverlappingPackageGroupType test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
	@if ! echo "${$@_output}" | grep -q "Could not create metadata for type: .*/test_input/negative/types.OverlappingPackageGroupType."; then \
	    echo "types.OverlappingPackageGroupType test failed"; \
	    echo "${$@_output}"; \
	    false; \
	fi
