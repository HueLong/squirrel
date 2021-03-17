<?php
/**
 * Created by PhpStorm.
 * User: bob
 * Date: 2020/5/15
 * Time: 10:55
 */

namespace App\Providers;

use Illuminate\Support\ServiceProvider;

class RepositoryServiceProvider extends ServiceProvider
{
    public function register()
    {
        $this->app->bind(
            'App\Repository\Api\ImageRepository'
        );
    }
}
