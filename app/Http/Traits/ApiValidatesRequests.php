<?php
/**
 * Created by PhpStorm.
 * User: bob
 * Date: 2020/5/27
 * Time: 16:33
 */

namespace App\Http\Traits;


use App\Constants\ExceptionCodeConst;
use Illuminate\Contracts\Validation\Factory;
use Illuminate\Contracts\Validation\Validator;
use Illuminate\Http\Exceptions\HttpResponseException;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;

trait ApiValidatesRequests
{
    public function apiValidate(Request $request, array $rules, array $messages = [], array $customAttributes = [])
    {
        $validator = $this->getValidationFactory()->make($request->all(), $rules, $messages, $customAttributes);
        if ($validator->fails()) {
            $this->throwApiValidationException($validator, ExceptionCodeConst::PARAMS_VALIDATE_ERROR);
        }
    }


    /**
     * Get a validation factory instance.
     *
     * @return \Illuminate\Contracts\Validation\Factory
     */
    protected function getValidationFactory()
    {
        return app(Factory::class);
    }

    protected function throwApiValidationException(Validator $validator, $code, $httpCode = 200)
    {
        throw new HttpResponseException(new JsonResponse([
            'messages' => $validator->errors()->getMessages(),
            'message'  => implode(',', $validator->errors()->getMessages()[array_key_first($validator->errors()->getMessages())]),
            'code'     => $code,
        ], $httpCode));
    }

    /**
     * 验证数组
     * @param array $array
     * @param array $rules
     * @param array $messages
     * @param array $customAttributes
     */
    public function apiValidateArray(array $array, array $rules, array $messages = [], array $customAttributes = [])
    {
        $validator = $this->getValidationFactory()->make($array, $rules, $messages, $customAttributes);
        if ($validator->fails()) {
            $this->throwApiValidationException($validator, ExceptionCodeConst::PARAMS_VALIDATE_ERROR);
        }
    }
}
