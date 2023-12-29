function babi(){
    $.ajax({
        url : 'index_.php?act=daily',
        type : 'POST',
        data : {
            itemId			: 75,
            periodId		: 3,
            selserver		: 2,
        },
        dataType : 'json',
        success : function(response){
            if (response.message == 'success'){
                alertMsg(response.data, 'success');
                setTimeout(function(){location.reload()}, 3000);
            }else{
                if (response.needLogin == 1){
                    $('.loginMethod').click();
                }
                alertMsg(response.data, 'danger');
                setTimeout(function() {
                    $this.button('reset');
                    form.data('submitted', false);
                }, 50000);
            }
        },
        error : function(data){
            alertMsg('Problem Connecting, Please Try Again', 'danger');
            setTimeout(function() {
                $this.button('reset');
                form.data('submitted', false);
            }, 50000);
        }
    });
}

babi();