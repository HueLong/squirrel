<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;

use Illuminate\Database\Eloquent\Model;

class PicGallery extends Model
{
	use HasDateTimeFormatter;
    protected $table = 'pic_gallery';


    public function setUpdatedAtAttribute($value)
    {
        $this->attributes['updated_at'] = $value ? strtotime($value) : 0;
    }

    public function setCreatedAtAttribute($value)
    {
        $this->attributes['created_at'] = $value ? strtotime($value) : 0;
    }
}
