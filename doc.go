// swagger:operation GET /examples Examples getExamples
//
// Retrieve a list of examples
//
// ---
// parameters:
// - $ref: '#/parameters/Authorization'
// responses:
//   200:
//     description: List of Examples
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/Example'
//   default:
//     $ref: '#/definitions/RespondError'

// swagger:operation GET /examples/{exampleID} Examples getExample
//
// Retrieve a specific example
//
// ---
// parameters:
// - $ref: '#/parameters/Authorization'
// - $ref: '#/parameters/ExampleIDParam'
// responses:
//   200:
//     description:
//     schema:
//       $ref: '#/definitions/Example'
//   default:
//     $ref: '#/definitions/RespondError'

// swagger:operation POST /examples Examples postExample
//
// Create an Example
//
// ---
// parameters:
// - $ref: '#/parameters/Authorization'
// - $ref: '#/parameters/PostEmployeeRequestPayload'
// responses:
//   '201':
//     description: Example Created Successfully
//   default:
//     $ref: '#/definitions/RespondError'

// swagger:operation PATCH /examples/{exampleID} Examples patchExample
//
// Update an Example
//
// ---
// parameters:
// - $ref: '#/parameters/Authorization'
// - $ref: '#/parameters/ExampleIDParam'
// - $ref: '#/parameters/PatchExampleRequestPayload'
// responses:
//   200:
//     description: Patched Example
//     schema:
//       type: array
//       items:
//         $ref: '#/definitions/Example'
//   default:
//     $ref: '#/definitions/RespondError'
package main
