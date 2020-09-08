# rstack

RStack is a recursive stack, meaning that every node in the stack is itself an
RStack.  There is no special "stack" struct with a pointer to the first (or
last) "node" struct in the stack.  The RStack you are pointing to _is_ the last
node in the list, and points to another RStack node, which is the last node in
_its_ list.

This is convenient for passing a list (stack) of results into a function, which
pushes a result and passes the cumulative results to another function (maybe
recursively).  Each function sees its caller's result list (stack), without
having to manage a global stack at each invocation.
