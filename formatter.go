// to define the Formatter interface for formatting
// based on different format
package naivelog

type Formatter interface {
	Format(entry *Entry) error
}
