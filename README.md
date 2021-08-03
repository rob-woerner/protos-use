The repos https://github.com/rob-woerner/protos and https://github.com/rob-woerner/protos-use provide an example solution to the problem of sharing protobuf definitions across files, packages, and modules.

I tried various tricks before finding this one.  The keys were to put the proto files in the root directory and set the output path (--go_out and --go-grpc_out) to ".".&nbsp;&nbsp;&nbsp;I don't think this should be necessary, but <b>many</b> attempts failed on some sort of package name mismatch.

If you find a solution that allows you to put the proto files in a subdirectory, I'd love to see it. 

To be a successful solution, you need to successfully run in protos:<br/>
&nbsp;&nbsp;&nbsp;make<br/>
&nbsp;&nbsp;&nbsp;go mod tidy<br/>
&nbsp;&nbsp;&nbsp;go build ./...<br/>

and run in protos-use:<br/>
&nbsp;&nbsp;&nbsp;go mod tidy<br/>
&nbsp;&nbsp;&nbsp;go build ./...<br/>


Tip o' the hat to stackoverflow user DazWilkin for invaluable suggestions.


