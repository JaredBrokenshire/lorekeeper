import {extend} from 'vee-validate'
import {min as rule_min} from 'vee-validate/dist/rules'

// eslint-disable-next-line object-curly-newline
import {validatorIsImageFile} from "./validators";

// ////////////////////////////////////////////////////////
// General
// ////////////////////////////////////////////////////////

export const min = extend('min', rule_min)

export const isImageFile = extend('isImageFile', {
  validate: validatorIsImageFile,
  message: 'Only JPG, JPEG, PNG and WEBP files are supported.'
})