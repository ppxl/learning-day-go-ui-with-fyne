ARTEFACT_ID=com.github.learningday2024
VERSION=0.1.0
WORKDIR:=$(shell pwd)
TARGET_DIR=${WORKDIR}/target
TMP_DIR:=$(BUILD_DIR)/tmp

${TARGET_DIR}:
	@mkdir -p $@

build-linux: ${TARGET_DIR}
	@cd ${TARGET_DIR} && fyne package -os linux -icon ${WORKDIR}/resources/avatar.png --sourceDir ${WORKDIR}

build-android: ${TARGET_DIR}
	@fyne package -os android -appID ${ARTEFACT_ID} -icon resources/avatar.png

install-android: ${TARGET_DIR}
	@fyne install -os android -appID ${ARTEFACT_ID} -icon resources/avatar.png