<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\PicGallery;
use App\Enum\GrantEnum;
use App\Service\ActivityRelationService;
use App\Service\CaseService;
use App\Service\GrantCateSettingService;
use App\Service\GrantService;
use App\Service\GrantTagService;
use App\Service\PicBankService;
use App\Service\PicGalleryService;
use App\Service\TagService;
use App\Service\UserService;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class PicGalleryController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new PicGallery(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('name');
            $grid->column('desc')->limit(8);
            $grid->column('pic_url')->image();;
            $grid->column('status');
            $grid->column('created_at');
            $grid->column('updated_at')->sortable();

            $grid->filter(function (Grid\Filter $filter) {
                $filter->equal('id');

            });
        });
    }

    /**
     * Make a show builder.
     *
     * @param mixed $id
     *
     * @return Show
     */
    protected function detail($id)
    {
        return Show::make($id, new PicGallery(), function (Show $show) {
            $show->field('id');
            $show->field('name');
            $show->field('desc');
            $show->field('pic_url');
            $show->field('status');
            $show->field('created_at');
            $show->field('updated_at');
        });
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        return Form::make(new PicGallery(), function (Form $form) {
            $form->disableResetButton();

            $form->display('id');
            $form->text('name');
            $form->text('desc');
            $form->multipleImage('pic_url_list', '作品图片')
                ->dir('squirrel/work')
                ->accept('jpg,png,jpeg', 'image/*')
                ->uniqueName()
                ->autoUpload()
                ->autoSave(false)
                ->removable(false)
                ->retainable()
                ->required();

            $form->saving(function (Form $form) {
                $input = $form->input();
                //原图和封面图处理
                $originPicUrl = explode(',', $input['pic_url_list']);
                $res = (new PicGalleryService())->batchInsert($input['name'], $input['desc'], $originPicUrl);
                if ($res) {
                    return $form->response()->success('保存成功')->redirect('pic_gallery');
                } else {
                    return $form->response()->error('系统错误')->redirect('pic_gallery');
                }
            });
        });
    }
}
