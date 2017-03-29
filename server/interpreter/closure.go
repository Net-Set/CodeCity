/* Copyright 2017 Google Inc.
 * https://github.com/NeilFraser/CodeCity
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package interpreter

import (
	"CodeCity/server/interpreter/ast"
	"CodeCity/server/interpreter/object"
)

// A closure is an object that can be called / applied.
type closure struct {
	object.Object
	scope  *scope
	params []string
	body   *ast.BlockStatement
}

// *Function must satisfy Value.
var _ object.Value = (*closure)(nil)

func (closure) ToString() object.String {
	return "[object Function]"
}

// newClosure returns a new closure object with the specified owner,
// scope and body, having parent functionProto.
func newClosure(owner *object.Owner, scope *scope,
	params []*ast.Identifier, body *ast.BlockStatement) *closure {
	var cl = new(closure)
	cl.Object = *object.New(owner, functionProto)
	cl.scope = scope
	cl.SetProperty("length", object.Number(len(params)))
	cl.params = make([]string, len(params))
	for i, p := range params {
		cl.params[i] = p.Name
	}
	cl.body = body
	return cl
}

// functionProto is the the (plain) JavaScript object that is the
// prototype for all closures.  (It would usually be accessed in
// JavaScript as Function.prototype.)
var functionProto = object.New(nil, object.ObjectProto)