<?php
/**
 * Created by PhpStorm.
 * User: bob
 * Date: 2020/6/27
 * Time: 01:46
 */

namespace App\Constants;


use App\Http\Common\WxPay\Log;
use Illuminate\Database\Eloquent\Builder;
use Spatie\Permission\Models\Permission;

class MenuVars
{
    public static function getMenu()
    {
       // \Log::info('----获取菜单start---');
//        $menus = \Cache::remember("admin_menus", 60*24, function () {
//            return \App\Models\Scratch\Permission::with([
//                'childs' => function ($query) {
//                    $query->with('icon');
//                },
//                'icon'
//            ])->where('parent_id', 0)->where('hash_id', '')->orderBy('sort', 'desc')->get();
//        });

        $menus = \App\Models\Scratch\Permission::with([
            'childs' => function ($query) {
                $query->with(['icon', 'childs'])->orderByRaw('sort desc, id');
            },
            'icon'
        ])->where('parent_id', 0)->orderByRaw('sort desc, id ')->get();

        //\Log::info('----获取菜单end---');

        return $menus;
    }
}
