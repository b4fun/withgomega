// Package withgomega generated by withgomega/gen. DO NOT EDIT.
// source: https://raw.githubusercontent.com/onsi/gomega/v1.9.0/matchers.go
package withgomega

import (
	"time"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

type Matcher struct{}

//Equal uses reflect.DeepEqual to compare actual with expected.  Equal is strict about
//types when performing comparisons.
//It is an error for both actual and expected to be nil.  Use BeNil() instead.
func (Matcher) Equal(expected interface{}) types.GomegaMatcher {
	return Equal(expected)
}

//BeEquivalentTo is more lax than Equal, allowing equality between different types.
//This is done by converting actual to have the type of expected before
//attempting equality with reflect.DeepEqual.
//It is an error for actual and expected to be nil.  Use BeNil() instead.
func (Matcher) BeEquivalentTo(expected interface{}) types.GomegaMatcher {
	return BeEquivalentTo(expected)
}

//BeIdenticalTo uses the == operator to compare actual with expected.
//BeIdenticalTo is strict about types when performing comparisons.
//It is an error for both actual and expected to be nil.  Use BeNil() instead.
func (Matcher) BeIdenticalTo(expected interface{}) types.GomegaMatcher {
	return BeIdenticalTo(expected)
}

//BeNil succeeds if actual is nil
func (Matcher) BeNil() types.GomegaMatcher {
	return BeNil()
}

//BeTrue succeeds if actual is true
func (Matcher) BeTrue() types.GomegaMatcher {
	return BeTrue()
}

//BeFalse succeeds if actual is false
func (Matcher) BeFalse() types.GomegaMatcher {
	return BeFalse()
}

//HaveOccurred succeeds if actual is a non-nil error
//The typical Go error checking pattern looks like:
//    err := SomethingThatMightFail()
//    Expect(err).ShouldNot(HaveOccurred())
func (Matcher) HaveOccurred() types.GomegaMatcher {
	return HaveOccurred()
}

//Succeed passes if actual is a nil error
//Succeed is intended to be used with functions that return a single error value. Instead of
//    err := SomethingThatMightFail()
//    Expect(err).ShouldNot(HaveOccurred())
//
//You can write:
//    Expect(SomethingThatMightFail()).Should(Succeed())
//
//It is a mistake to use Succeed with a function that has multiple return values.  Gomega's ?? and Expect
//functions automatically trigger failure if any return values after the first return value are non-zero/non-nil.
//This means that ??(MultiReturnFunc()).ShouldNot(Succeed()) can never pass.
func (Matcher) Succeed() types.GomegaMatcher {
	return Succeed()
}

//MatchError succeeds if actual is a non-nil error that matches the passed in string/error.
//
//These are valid use-cases:
//  Expect(err).Should(MatchError("an error")) //asserts that err.Error() == "an error"
//  Expect(err).Should(MatchError(SomeError)) //asserts that err == SomeError (via reflect.DeepEqual)
//
//It is an error for err to be nil or an object that does not implement the Error interface
func (Matcher) MatchError(expected interface{}) types.GomegaMatcher {
	return MatchError(expected)
}

//BeClosed succeeds if actual is a closed channel.
//It is an error to pass a non-channel to BeClosed, it is also an error to pass nil
//
//In order to check whether or not the channel is closed, Gomega must try to read from the channel
//(even in the `ShouldNot(BeClosed())` case).  You should keep this in mind if you wish to make subsequent assertions about
//values coming down the channel.
//
//Also, if you are testing that a *buffered* channel is closed you must first read all values out of the channel before
//asserting that it is closed (it is not possible to detect that a buffered-channel has been closed until all its buffered values are read).
//
//Finally, as a corollary: it is an error to check whether or not a send-only channel is closed.
func (Matcher) BeClosed() types.GomegaMatcher {
	return BeClosed()
}

//Receive succeeds if there is a value to be received on actual.
//Actual must be a channel (and cannot be a send-only channel) -- anything else is an error.
//
//Receive returns immediately and never blocks:
//
//- If there is nothing on the channel `c` then Expect(c).Should(Receive()) will fail and ??(c).ShouldNot(Receive()) will pass.
//
//- If the channel `c` is closed then Expect(c).Should(Receive()) will fail and ??(c).ShouldNot(Receive()) will pass.
//
//- If there is something on the channel `c` ready to be read, then Expect(c).Should(Receive()) will pass and ??(c).ShouldNot(Receive()) will fail.
//
//If you have a go-routine running in the background that will write to channel `c` you can:
//    Eventually(c).Should(Receive())
//
//This will timeout if nothing gets sent to `c` (you can modify the timeout interval as you normally do with `Eventually`)
//
//A similar use-case is to assert that no go-routine writes to a channel (for a period of time).  You can do this with `Consistently`:
//    Consistently(c).ShouldNot(Receive())
//
//You can pass `Receive` a matcher.  If you do so, it will match the received object against the matcher.  For example:
//    Expect(c).Should(Receive(Equal("foo")))
//
//When given a matcher, `Receive` will always fail if there is nothing to be received on the channel.
//
//Passing Receive a matcher is especially useful when paired with Eventually:
//
//    Eventually(c).Should(Receive(ContainSubstring("bar")))
//
//will repeatedly attempt to pull values out of `c` until a value matching "bar" is received.
//
//Finally, if you want to have a reference to the value *sent* to the channel you can pass the `Receive` matcher a pointer to a variable of the appropriate type:
//    var myThing thing
//    Eventually(thingChan).Should(Receive(&myThing))
//    Expect(myThing.Sprocket).Should(Equal("foo"))
//    Expect(myThing.IsValid()).Should(BeTrue())
func (Matcher) Receive(args ...interface{}) types.GomegaMatcher {
	return Receive(args...)
}

//BeSent succeeds if a value can be sent to actual.
//Actual must be a channel (and cannot be a receive-only channel) that can sent the type of the value passed into BeSent -- anything else is an error.
//In addition, actual must not be closed.
//
//BeSent never blocks:
//
//- If the channel `c` is not ready to receive then Expect(c).Should(BeSent("foo")) will fail immediately
//- If the channel `c` is eventually ready to receive then Eventually(c).Should(BeSent("foo")) will succeed.. presuming the channel becomes ready to receive  before Eventually's timeout
//- If the channel `c` is closed then Expect(c).Should(BeSent("foo")) and ??(c).ShouldNot(BeSent("foo")) will both fail immediately
//
//Of course, the value is actually sent to the channel.  The point of `BeSent` is less to make an assertion about the availability of the channel (which is typically an implementation detail that your test should not be concerned with).
//Rather, the point of `BeSent` is to make it possible to easily and expressively write tests that can timeout on blocked channel sends.
func (Matcher) BeSent(arg interface{}) types.GomegaMatcher {
	return BeSent(arg)
}

//MatchRegexp succeeds if actual is a string or stringer that matches the
//passed-in regexp.  Optional arguments can be provided to construct a regexp
//via fmt.Sprintf().
func (Matcher) MatchRegexp(regexp string, args ...interface{}) types.GomegaMatcher {
	return MatchRegexp(regexp, args...)
}

//ContainSubstring succeeds if actual is a string or stringer that contains the
//passed-in substring.  Optional arguments can be provided to construct the substring
//via fmt.Sprintf().
func (Matcher) ContainSubstring(substr string, args ...interface{}) types.GomegaMatcher {
	return ContainSubstring(substr, args...)
}

//HavePrefix succeeds if actual is a string or stringer that contains the
//passed-in string as a prefix.  Optional arguments can be provided to construct
//via fmt.Sprintf().
func (Matcher) HavePrefix(prefix string, args ...interface{}) types.GomegaMatcher {
	return HavePrefix(prefix, args...)
}

//HaveSuffix succeeds if actual is a string or stringer that contains the
//passed-in string as a suffix.  Optional arguments can be provided to construct
//via fmt.Sprintf().
func (Matcher) HaveSuffix(suffix string, args ...interface{}) types.GomegaMatcher {
	return HaveSuffix(suffix, args...)
}

//MatchJSON succeeds if actual is a string or stringer of JSON that matches
//the expected JSON.  The JSONs are decoded and the resulting objects are compared via
//reflect.DeepEqual so things like key-ordering and whitespace shouldn't matter.
func (Matcher) MatchJSON(json interface{}) types.GomegaMatcher {
	return MatchJSON(json)
}

//MatchXML succeeds if actual is a string or stringer of XML that matches
//the expected XML.  The XMLs are decoded and the resulting objects are compared via
//reflect.DeepEqual so things like whitespaces shouldn't matter.
func (Matcher) MatchXML(xml interface{}) types.GomegaMatcher {
	return MatchXML(xml)
}

//MatchYAML succeeds if actual is a string or stringer of YAML that matches
//the expected YAML.  The YAML's are decoded and the resulting objects are compared via
//reflect.DeepEqual so things like key-ordering and whitespace shouldn't matter.
func (Matcher) MatchYAML(yaml interface{}) types.GomegaMatcher {
	return MatchYAML(yaml)
}

//BeEmpty succeeds if actual is empty.  Actual must be of type string, array, map, chan, or slice.
func (Matcher) BeEmpty() types.GomegaMatcher {
	return BeEmpty()
}

//HaveLen succeeds if actual has the passed-in length.  Actual must be of type string, array, map, chan, or slice.
func (Matcher) HaveLen(count int) types.GomegaMatcher {
	return HaveLen(count)
}

//HaveCap succeeds if actual has the passed-in capacity.  Actual must be of type array, chan, or slice.
func (Matcher) HaveCap(count int) types.GomegaMatcher {
	return HaveCap(count)
}

//BeZero succeeds if actual is the zero value for its type or if actual is nil.
func (Matcher) BeZero() types.GomegaMatcher {
	return BeZero()
}

//ContainElement succeeds if actual contains the passed in element.
//By default ContainElement() uses Equal() to perform the match, however a
//matcher can be passed in instead:
//    Expect([]string{"Foo", "FooBar"}).Should(ContainElement(ContainSubstring("Bar")))
//
//Actual must be an array, slice or map.
//For maps, ContainElement searches through the map's values.
func (Matcher) ContainElement(element interface{}) types.GomegaMatcher {
	return ContainElement(element)
}

//BeElementOf succeeds if actual is contained in the passed in elements.
//BeElementOf() always uses Equal() to perform the match.
//When the passed in elements are comprised of a single element that is either an Array or Slice, BeElementOf() behaves
//as the reverse of ContainElement() that operates with Equal() to perform the match.
//    Expect(2).Should(BeElementOf([]int{1, 2}))
//    Expect(2).Should(BeElementOf([2]int{1, 2}))
//Otherwise, BeElementOf() provides a syntactic sugar for Or(Equal(_), Equal(_), ...):
//    Expect(2).Should(BeElementOf(1, 2))
//
//Actual must be typed.
func (Matcher) BeElementOf(elements ...interface{}) types.GomegaMatcher {
	return BeElementOf(elements...)
}

//ConsistOf succeeds if actual contains precisely the elements passed into the matcher.  The ordering of the elements does not matter.
//By default ConsistOf() uses Equal() to match the elements, however custom matchers can be passed in instead.  Here are some examples:
//
//    Expect([]string{"Foo", "FooBar"}).Should(ConsistOf("FooBar", "Foo"))
//    Expect([]string{"Foo", "FooBar"}).Should(ConsistOf(ContainSubstring("Bar"), "Foo"))
//    Expect([]string{"Foo", "FooBar"}).Should(ConsistOf(ContainSubstring("Foo"), ContainSubstring("Foo")))
//
//Actual must be an array, slice or map.  For maps, ConsistOf matches against the map's values.
//
//You typically pass variadic arguments to ConsistOf (as in the examples above).  However, if you need to pass in a slice you can provided that it
//is the only element passed in to ConsistOf:
//
//    Expect([]string{"Foo", "FooBar"}).Should(ConsistOf([]string{"FooBar", "Foo"}))
//
//Note that Go's type system does not allow you to write this as ConsistOf([]string{"FooBar", "Foo"}...) as []string and []interface{} are different types - hence the need for this special rule.
func (Matcher) ConsistOf(elements ...interface{}) types.GomegaMatcher {
	return ConsistOf(elements...)
}

//ContainElements succeeds if actual contains the passed in elements. The ordering of the elements does not matter.
//By default ContainElements() uses Equal() to match the elements, however custom matchers can be passed in instead. Here are some examples:
//
//    Expect([]string{"Foo", "FooBar"}).Should(ContainElements("FooBar"))
//    Expect([]string{"Foo", "FooBar"}).Should(ContainElements(ContainSubstring("Bar"), "Foo"))
//
//Actual must be an array, slice or map.
//For maps, ContainElements searches through the map's values.
func (Matcher) ContainElements(elements ...interface{}) types.GomegaMatcher {
	return ContainElements(elements...)
}

//HaveKey succeeds if actual is a map with the passed in key.
//By default HaveKey uses Equal() to perform the match, however a
//matcher can be passed in instead:
//    Expect(map[string]string{"Foo": "Bar", "BazFoo": "Duck"}).Should(HaveKey(MatchRegexp(`.+Foo$`)))
func (Matcher) HaveKey(key interface{}) types.GomegaMatcher {
	return HaveKey(key)
}

//HaveKeyWithValue succeeds if actual is a map with the passed in key and value.
//By default HaveKeyWithValue uses Equal() to perform the match, however a
//matcher can be passed in instead:
//    Expect(map[string]string{"Foo": "Bar", "BazFoo": "Duck"}).Should(HaveKeyWithValue("Foo", "Bar"))
//    Expect(map[string]string{"Foo": "Bar", "BazFoo": "Duck"}).Should(HaveKeyWithValue(MatchRegexp(`.+Foo$`), "Bar"))
func (Matcher) HaveKeyWithValue(key interface{}, value interface{}) types.GomegaMatcher {
	return HaveKeyWithValue(key, value)
}

//BeNumerically performs numerical assertions in a type-agnostic way.
//Actual and expected should be numbers, though the specific type of
//number is irrelevant (float32, float64, uint8, etc...).
//
//There are six, self-explanatory, supported comparators:
//    Expect(1.0).Should(BeNumerically("==", 1))
//    Expect(1.0).Should(BeNumerically("~", 0.999, 0.01))
//    Expect(1.0).Should(BeNumerically(">", 0.9))
//    Expect(1.0).Should(BeNumerically(">=", 1.0))
//    Expect(1.0).Should(BeNumerically("<", 3))
//    Expect(1.0).Should(BeNumerically("<=", 1.0))
func (Matcher) BeNumerically(comparator string, compareTo ...interface{}) types.GomegaMatcher {
	return BeNumerically(comparator, compareTo...)
}

//BeTemporally compares time.Time's like BeNumerically
//Actual and expected must be time.Time. The comparators are the same as for BeNumerically
//    Expect(time.Now()).Should(BeTemporally(">", time.Time{}))
//    Expect(time.Now()).Should(BeTemporally("~", time.Now(), time.Second))
func (Matcher) BeTemporally(comparator string, compareTo time.Time, threshold ...time.Duration) types.GomegaMatcher {
	return BeTemporally(comparator, compareTo, threshold...)
}

//BeAssignableToTypeOf succeeds if actual is assignable to the type of expected.
//It will return an error when one of the values is nil.
//    Expect(0).Should(BeAssignableToTypeOf(0))         // Same values
//    Expect(5).Should(BeAssignableToTypeOf(-1))        // different values same type
//    Expect("foo").Should(BeAssignableToTypeOf("bar")) // different values same type
//    Expect(struct{ Foo string }{}).Should(BeAssignableToTypeOf(struct{ Foo string }{}))
func (Matcher) BeAssignableToTypeOf(expected interface{}) types.GomegaMatcher {
	return BeAssignableToTypeOf(expected)
}

//Panic succeeds if actual is a function that, when invoked, panics.
//Actual must be a function that takes no arguments and returns no results.
func (Matcher) Panic() types.GomegaMatcher {
	return Panic()
}

//BeAnExistingFile succeeds if a file exists.
//Actual must be a string representing the abs path to the file being checked.
func (Matcher) BeAnExistingFile() types.GomegaMatcher {
	return BeAnExistingFile()
}

//BeARegularFile succeeds if a file exists and is a regular file.
//Actual must be a string representing the abs path to the file being checked.
func (Matcher) BeARegularFile() types.GomegaMatcher {
	return BeARegularFile()
}

//BeADirectory succeeds if a file exists and is a directory.
//Actual must be a string representing the abs path to the file being checked.
func (Matcher) BeADirectory() types.GomegaMatcher {
	return BeADirectory()
}

//And succeeds only if all of the given matchers succeed.
//The matchers are tried in order, and will fail-fast if one doesn't succeed.
//  Expect("hi").To(And(HaveLen(2), Equal("hi"))
//
//And(), Or(), Not() and WithTransform() allow matchers to be composed into complex expressions.
func (Matcher) And(ms ...types.GomegaMatcher) types.GomegaMatcher {
	return And(ms...)
}

//SatisfyAll is an alias for And().
//  Expect("hi").Should(SatisfyAll(HaveLen(2), Equal("hi")))
func (Matcher) SatisfyAll(matchers ...types.GomegaMatcher) types.GomegaMatcher {
	return SatisfyAll(matchers...)
}

//Or succeeds if any of the given matchers succeed.
//The matchers are tried in order and will return immediately upon the first successful match.
//  Expect("hi").To(Or(HaveLen(3), HaveLen(2))
//
//And(), Or(), Not() and WithTransform() allow matchers to be composed into complex expressions.
func (Matcher) Or(ms ...types.GomegaMatcher) types.GomegaMatcher {
	return Or(ms...)
}

//SatisfyAny is an alias for Or().
//  Expect("hi").SatisfyAny(Or(HaveLen(3), HaveLen(2))
func (Matcher) SatisfyAny(matchers ...types.GomegaMatcher) types.GomegaMatcher {
	return SatisfyAny(matchers...)
}

//Not negates the given matcher; it succeeds if the given matcher fails.
//  Expect(1).To(Not(Equal(2))
//
//And(), Or(), Not() and WithTransform() allow matchers to be composed into complex expressions.
func (Matcher) Not(matcher types.GomegaMatcher) types.GomegaMatcher {
	return Not(matcher)
}

//WithTransform applies the `transform` to the actual value and matches it against `matcher`.
//The given transform must be a function of one parameter that returns one value.
//  var plus1 = func(i int) int { return i + 1 }
//  Expect(1).To(WithTransform(plus1, Equal(2))
//
//And(), Or(), Not() and WithTransform() allow matchers to be composed into complex expressions.
func (Matcher) WithTransform(transform interface{}, matcher types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(transform, matcher)
}
