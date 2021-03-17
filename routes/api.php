<?php

use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::group(['namespace' => 'Oss'], function () {
    Route::get('/oss_policy', 'AliYunOssSignController@index');
});

Route::group(['namespace' => 'Web'], function () {
    Route::get('/list_img', 'ImageController@listImage');
    Route::post('/upload_img', 'ImageController@saveImage');
});

