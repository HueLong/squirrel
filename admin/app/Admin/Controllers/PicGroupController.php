<?php

namespace App\Admin\Controllers;

use App\Admin\Repositories\PicGroup;
use App\Service\PicGalleryService;
use Dcat\Admin\Form;
use Dcat\Admin\Grid;
use Dcat\Admin\Show;
use Dcat\Admin\Http\Controllers\AdminController;

class PicGroupController extends AdminController
{
    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        return Grid::make(new PicGroup(), function (Grid $grid) {
            $grid->column('id')->sortable();
            $grid->column('name');
            $grid->column('cover')->image();
            $grid->column('desc');
            $grid->column('status');
            $grid->column('created_at');

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
        return Show::make($id, new PicGroup(), function (Show $show) {
            $show->field('id');
            $show->field('name');
            $show->field('cover');
            $show->field('desc');
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
        return Form::make(new PicGroup(), function (Form $form) {
            $form->disableResetButton();

            $form->display('id');
            $form->text('name');
            $form->text('desc');
            $form->image('cover', '封面图')
                ->dir('squirrel/work')
                ->accept('jpg,png,jpeg', 'image/*')
                ->uniqueName()
                ->autoUpload()
                ->autoSave(false)
                ->removable(false)
                ->retainable()
                ->required();
        });
    }
}
