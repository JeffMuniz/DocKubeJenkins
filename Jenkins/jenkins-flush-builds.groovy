

def jobName = "EC/est-api-autoec/develop"
def job = Jenkins.getInstance().getItemByFullName(jobName, Job.class)
job.getBuilds().each { it.delete() }
job.nextBuildNumber = 1
job.save()

def jobName = "EC/est-api-autoec/staging"
def job = Jenkins.getInstance().getItemByFullName(jobName, Job.class)
job.getBuilds().each { it.delete() }
job.nextBuildNumber = 1
job.save()

def jobName = "EC/est-api-autoec/master"
def job = Jenkins.getInstance().getItemByFullName(jobName, Job.class)
job.getBuilds().each { it.delete() }
job.nextBuildNumber = 1
job.save()


https://jenkins-qa.maquiina.com/script



