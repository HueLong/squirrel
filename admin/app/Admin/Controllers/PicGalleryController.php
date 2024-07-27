<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\PicGallery;
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
            $form->display('id');
            $form->text('name');
            $form->text('desc');
            $form->image('pic_url', '作品图片')
                ->dir('squirrel/work')
                ->accept('jpg,png,jpeg', 'image/*')
                ->uniqueName()
                ->autoUpload()
                ->autoSave(false)
                ->removable(false)
                ->retainable()
//                ->saving(function ($value) {
//                    return $value ? get_img_link($value, 'free') : '';
//                })
                ->required();
        });
    }
}
