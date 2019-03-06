## Prepare debug env

- Refer to this in addition to AWS docs: https://medium.com/@csmadhav/awesomeness-of-golang-aws-lambda-d6bd08131117
- AWS docs: https://docs.aws.amazon.com/ja_jp/serverless-application-model/latest/developerguide/serverless-sam-cli-using-debugging-golang.html
- Install SAM via homebrew (added to brew bundle file)
- Add a AWS SAM Template config file

```yml
Resources:
  tars:
    Type: AWS::Serverless::Function
    Properties:
      Handler: bin/tars
      Runtime: go1.x
```

- Build dlv for Linux: `export DLV_PATH=$GOPATH/bin/linux/dlv; GOARCH=amd64 GOOS=linux go build -o $DLV_PATH github.com/derekparker/delve/cmd/dlv`
- Invoke SAM: `sam local start-api -d 5986 --debugger-path $DLV_PATH`
- Build in debug mode: `GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o <output path> <path to code directory>`
- Attach to debug session w/ Vim (sebdah/vim-delve) or VSCode
