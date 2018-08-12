# What is it?

[![Build Status](https://travis-ci.org/olivere/sqltype.svg?branch=master)](https://travis-ci.org/olivere/sqltype)

The `sqltype` package is a set of types that embraces the
`sql.Scanner` interface from `database/sql` to other types,
e.g. `time.Time` with `sqltype.NullTime` and `time.Duration`
with `sqltype.NullDuration`.

# Prior art

Parts of this are taken from the `lib/pq` library, e.g.
the `sqltype.NullTime` has been take from
[here](https://github.com/lib/pq/blob/8c6ee72f3e6bcb1542298dd5f76cb74af9742cec/encode.go#L583-L603).

# License

MIT. See [LICENSE](https://github.com/olivere/sqltype/blob/master/LICENSE) file.
