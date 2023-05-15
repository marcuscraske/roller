DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" > /dev/null && pwd )"

ARCH=(
  "darwin,amd64,roller"
  "darwin,arm64,roller"
  "linux,amd64,roller"
  "linux,arm,roller"
  "linux,arm64,roller"
  "windows,amd64,roller.exe"
)

echo "Starting..."

# Delete existing bin dir
if [[ -d ${DIR}/bin ]]; then
  rm -r ${DIR}/bin
  echo "Deleted ${DIR}/bin"
fi

BUILD_VERSION=$(git describe --tags --abbrev=0)
BUILD_COMMIT=$(git rev-parse --short HEAD)
BUILD_DATE=$(date +"%Y-%m-%d %H:%M:%S")

echo "Build Version: ${BUILD_VERSION}"
echo "Build Commit:  ${BUILD_COMMIT}"
echo "Build Date:    ${BUILD_DATE}"

# Build each architecture
for key in ${ARCH[@]}; do
  IFS=',' read -ra parts <<< "$key"
  echo "Building ${parts[0]}-${parts[1]}..."
  GOOS=${parts[0]} GOARCH=${parts[1]} \
    go build \
    -ldflags \
    " \
      -X 'roller/pkg/command.BuildVersion=${BUILD_VERSION}' \
      -X 'roller/pkg/command.BuildCommit=${BUILD_COMMIT}' \
      -X 'roller/pkg/command.BuildDate=${BUILD_DATE}' \
    " \
    -o ${DIR}/bin/${parts[0]}-${parts[1]}/${parts[2]} \
    cmd/main.go

  ( \
    cd ${DIR}/bin && \
    zip -j ${parts[0]}-${parts[1]}-${BUILD_VERSION}.zip ${DIR}/bin/${parts[0]}-${parts[1]}/${parts[2]} \
  )
done

echo "Finished!"
